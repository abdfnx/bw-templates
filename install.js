import { createWriteStream } from "fs";
import * as fs from "fs/promises";
import fetch from "node-fetch";
import { pipeline } from "stream/promises";
import StreamZip from "node-stream-zip";
import { ARCH_MAPPING, CONFIG, PLATFORM_MAPPING } from "./npm-config.js";
import path from "path";

const install = async () => {
  const packageJson = await fs.readFile("package.json").then(JSON.parse);

  let version = packageJson.version;

  if (typeof version !== "string") {
    throw new Error("Missing version in package.json");
  }

  if (version[0] === "v") version = version.slice(1);

  // Fetch Static Config
  let { name: binName, path: binPath, url } = CONFIG;

  url = url.replace(/{{arch}}/g, ARCH_MAPPING[process.arch]);
  url = url.replace(/{{platform}}/g, PLATFORM_MAPPING[process.platform]);
  url = url.replace(/{{version}}/g, version);
  url = url.replace(/{{bin_name}}/g, binName);

  let execType = "";

  if (process.platform == "win32") {
    execType = ".exe";
  }

  const folder = (old) => {
    let f =
      binName +
      "_" +
      PLATFORM_MAPPING[process.platform] +
      "_" +
      "v" +
      version +
      "_" +
      ARCH_MAPPING[process.arch];

    if (old == "yes") {
      return path.join(f, "bin", `create-botway-bot${execType}`);
    } else if (old == "no") {
      return path.join("bin", `create-botway-bot${execType}`);
    } else {
      return f;
    }
  };

  const response = await fetch(url);

  if (!response.ok) {
    throw new Error("Failed fetching the binary: " + response.statusText);
  }

  const zipFile = "create-botway-bot.zip";

  await pipeline(response.body, createWriteStream(zipFile));

  const zip = new StreamZip.async({ file: zipFile });

  const count = await zip.extract(null, ".");

  console.log(`Extracted ${count} entries`);

  await zip.close();

  if (process.platform != "win32") {
    await fs.rename(folder("yes"), folder("no"), (err) => {
      if (err) throw err;
    });

    await fs.rm(folder(), { recursive: true });
  }

  // chmod +x /bin/create-botway-bot
  await fs.chmod(path.join("bin", `create-botway-bot${execType}`), 0o755);

  await fs.rm(zipFile);
};

install()
  .then(async () => {
    process.exit(0);
  })
  .catch(async (err) => {
    console.error(err);
    process.exit(1);
  });

#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');

const binName = {
  win32: {
    x64: 'StylelintBC_windows_amd64_v1/stylelintbc.exe',
    arm64: 'StylelintBC_windows_arm64/stylelintbc.exe',
  },
  darwin: {
    x64: 'StylelintBC_darwin_amd64_v1/stylelintbc',
    arm64: 'StylelintBC_darwin_arm64/stylelintbc',
  },
  linux: {
    x64: 'StylelintBC_linux_amd64_v1/StylelintBC',
    arm64: 'StylelintBC_linux_arm64/stylelintbc',
  },
};

const platform = os.platform();
const arch = os.arch();

// Verify if the platform is supported
if (!binName[platform]) {
  console.error(`Platform not supported: ${platform}`);
  process.exit(1);
}

// Verify if the architecture is supported
if (!binName[platform][arch]) {
  console.error(`Architecture ${arch} not supported for platform ${platform}`);
  process.exit(1);
}

const binPath = path.join(__dirname, 'bin', binName[platform][arch]);

// Process arguments to remove additional spaces
const args = process.argv.slice(2).map(arg => arg.trim());

const proc = spawn(binPath, args, { stdio: 'inherit' });

proc.on('close', (code) => {
  process.exit(code);
});
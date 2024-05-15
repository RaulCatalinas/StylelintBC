// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

// Third-Party libraries
import { exists as existsFolderOrFile } from 'fs-extra'

export async function exists(path: string) {
  try {
    return await existsFolderOrFile(path)
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('CheckFileExists')
    })

    process.exit(1)
  }
}

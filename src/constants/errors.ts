import { ISSUES } from './github'

export const ERROR_MESSAGES = {
  NotFound: "The package.json file wasn't found in the current directory.",

  Default: `Something went wrong, please try again later, if the error persists please report it on ${ISSUES}.`,

  Dependencies: `An error occurred while installing dependencies, please try again later, if the error persists please report it on ${ISSUES}.`,

  Stylelint: `An error has occurred during the Stylelint configuration process, please try again later, if the error persists please report it on ${ISSUES}.`,

  StylelintConfigCleanOrder: `An error has occurred during the stylelint-config-clean-order configuration process, please try again later, if the error persists please report it on ${ISSUES}.`,

  PackageManagerSelection: `An error occurred while selecting the package manager, please try again later, if the error persists, please report it on ${ISSUES}.`,

  CheckFileExists: `An error occurred while checking if the file/folder exists, please try again later, if the error persists, please report it on ${ISSUES}.`
} as const
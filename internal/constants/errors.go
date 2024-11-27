package constants

type ErrorMessage string

var ERROR_MESSAGES = map[string]ErrorMessage{
	"NotFound": "The package.json file wasn't found in the current directory.",

	"Default": "Something went wrong, please try again later, if the error persists please report it on " + ISSUES + ".",

	"Dependencies": "An error occurred while installing dependencies, please try again later, if the error persists please report it on " + ISSUES + ".",

	"Stylelint": "An error has occurred during the Stylelint configuration process, please try again later, if the error persists please report it on " + ISSUES + ".",

	"StylelintConfigCleanOrder": "An error has occurred during the stylelint-config-clean-order configuration process, please try again later, if the error persists please report it on " + ISSUES + ".",

	"PackageManagerSelection": "An error occurred while selecting the package manager, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"VSCodeEditor": "An error occurred while prompting the user about using VS Code as the code editor. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"CreateEmptyFile": "An error occurred while creating the empty file. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"VSCodeConfig": "An error occurred while configuring VS Code. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"AddRecommendedExtensions": "An error occurred while adding recommended extensions to the 'extensions.json' file. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"CreateFolder": "An error occurred while creating the folder: {folderName}, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"ConfigFilesCreate": "An error occurred while creating configuration files. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"PublishConfirmation": "An error occurred while confirming npm publication. Please try again later, if the error persists, please report it on" + ISSUES + ".",

	"NpmIgnoreWrite": "An error occurred while writing to the '.npmignore' file. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"GitHubRepoOpen": "An error occurred while opening the GitHub repository in a new browser tab. Please try again later, if the error persists please report it on" + ISSUES + ".",
}

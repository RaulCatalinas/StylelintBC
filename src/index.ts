// Config
import { configureCLI } from './config/cli'
import { configureDefaultOption, configureOptions } from './config/options'

// Constants
import { stylelintbc } from './constants/stylelintbc'

configureCLI()
configureOptions()

stylelintbc.parse()

configureDefaultOption()

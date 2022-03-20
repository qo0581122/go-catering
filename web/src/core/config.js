/**
 * 网站配置文件
 */

const config = {
  appName: 'Catering',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Catering，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:V2.5.0 beta.2`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信：shouzi_1994 QQ群：622360840`
      )
    )
    console.log('\n')
  }
}

export default config

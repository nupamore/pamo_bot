const request = require('request')
const CONFIG = require('./../config.json')

// aws
const AWS = require('aws-sdk')
AWS.config.region = 'ap-northeast-2'
AWS.config.credentials = new AWS.Credentials(
    CONFIG.aws.awsAccessKeyId,
    CONFIG.aws.awsSecretAccessKey,
)
const awsTranslate = new AWS.Translate({
    endpoint: new AWS.Endpoint(CONFIG.aws.endpoint),
    region: AWS.config.region,
})

/**
 * aws translate
 *
 * @param {Object} message
 * @param {String} source   before language
 * @param {String} target   after language
 */
function aws(message, source, target) {
    awsTranslate.translateText(
        {
            SourceLanguageCode: source,
            TargetLanguageCode: target,
            Text: message._.text.slice(target.length).trim(),
        },
        (err, data) => {
            if (!err) {
                message.channel.send(
                    `${message.author}(${data.SourceLanguageCode}): ${data.TranslatedText}`,
                )
            } else {
                message.channel.send(`errorCode: ${err.code}`)
            }
        },
    )
}

/**
 * Papago translate
 * api link: https://developers.naver.com/products/nmt/
 *
 * @param {Object} message
 * @param {String} type     nmt or smt
 * @param {String} source   before language
 * @param {String} target   after language
 */
function papago(message, type, source, target) {
    const options = {
        url: CONFIG.naver[type],
        form: { source, target, text: message._.text },
        headers: {
            'X-Naver-Client-Id': CONFIG.naver.id,
            'X-Naver-Client-Secret': CONFIG.naver.secret,
        },
    }
    request.post(options, (error, response, body) => {
        if (!error && response.statusCode == 200) {
            const translatedText = JSON.parse(body).message.result
                .translatedText
            message.channel.send(`${message.author}: ${translatedText}`)
        } else if (response.statusCode == 429) {
            message.channel.send(`Sorry, request limit is over (ㅠ_ㅠ)`)
        } else {
            message.channel.send(`errorCode: ${response.statusCode}`)
        }
    })
}

/**
 * Kakao translate
 * api link: https://developers.kakao.com/docs/restapi/translation
 *
 * @param {Object} message
 * @param {String} source   before language
 * @param {String} target   after language
 */
function kakao(message, source, target) {
    const options = {
        url: CONFIG.kakao.translate,
        form: { src_lang: source, target_lang: target, query: message._.text },
        headers: { Authorization: CONFIG.kakao.key },
    }
    request.post(options, (error, response, body) => {
        if (!error && response.statusCode == 200) {
            const translatedText = JSON.parse(body).translated_text
            message.channel.send(`${message.author}: ${translatedText}`)
        } else {
            message.channel.send(`errorCode: ${response.statusCode}`)
        }
    })
}

/**
 * Select api
 *
 * @param {Object} message
 * @param {String} type
 * @param {String} source   before language
 * @param {String} target   after language
 */
function translate(message, type, source, target) {
    const deprecatedMessage = `This command has been deprecated, Please use **${CONFIG.discord.prefix}t**`

    if (type.match('aws')) {
        aws(message, source, target)
    } else if (type.match(/nmt|smt/)) {
        // papago(message, type, source, target)
        message.channel.send(deprecatedMessage)
    } else if (type.match('kakao')) {
        // kakao(message, source, target)
        message.channel.send(deprecatedMessage)
    }
}

translate.comment = [
    'Translate',
    `**${CONFIG.discord.prefix}t en**
    Any language -> English
    https://docs.aws.amazon.com/translate/latest/dg/what-is.html#what-is-languages`,
]
module.exports = translate

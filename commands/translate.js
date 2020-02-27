const request = require('request')
const CONFIG = require('./../config.json')

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
    if (type.match(/nmt|smt/)) {
        papago(message, type, source, target)
    } else if (type.match('kakao')) {
        kakao(message, source, target)
    }
}

translate.comment =
    `**${CONFIG.discord.prefix}kj**` +
    ` - Korean -> Japanese (+ English, Chinese, Spanish, French)`
module.exports = translate

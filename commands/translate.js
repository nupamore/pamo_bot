
const request = require('request')
const CONFIG = require('./../config.json')

/**
 * Translate language
 * api link: https://developers.naver.com/products/nmt/
 * 
 * @param {Object} message
 * @param {String} type     nmt or smt
 * @param {String} source   before language
 * @param {String} target   after language
 */
function translate(message, type, source, target) {
    const options = {
        url: CONFIG.naver[type],
        form: { source, target, text: message._.text },
        headers: { 'X-Naver-Client-Id': CONFIG.naver.id, 'X-Naver-Client-Secret': CONFIG.naver.secret }
    }
    request.post(options, (error, response, body) => {
        if (!error && response.statusCode == 200) {
            const translatedText = JSON.parse(body).message.result.translatedText
            message.channel.send(`${ message.author }: ${ translatedText }`)
        }
        else if (response.statusCode == 429) {
            message.channel.send(`Sorry, request limit is over (ㅠ_ㅠ)`)
        }
        else {
            message.channel.send(`errorCode: ${ response.statusCode }`)
        }
    })
}


translate.comment = `!kj - Korean -> Japanese (+ English, Spanish, French)`
module.exports = translate;
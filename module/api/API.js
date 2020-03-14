module.exports = {
    LOGOUT: ['GET', '/auth/logout'],
    GET_PROFILE: ['GET', '/auth/me'],
    GET_UPLOADERS: ['GET', '/api/guild/{guildId}/uploaders'],
    GET_IMAGES: ['GET', '/api/guild/{guildId}/images'],
    DELETE_IMAGE: ['DELETE', '/api/image/{imageId}'],
}

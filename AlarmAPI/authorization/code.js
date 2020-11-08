const db = require("../db/db");
const jwt = require("jwt-decode");

module.exports = {
    /**
     * @param {string} code Code that refers to an auth token
     * @param {function(err, boolean)} next Function that is called when token has been validated
     */
    hasValidToken: (code, next) => {
        db.getTokenByCode(code, (err, token) => {
            if (err) {
                console.log(err);
                next(err, false);
                return;
            }

            if (!token || !token.token || !token.token.access_token) {
                next(err, false);
                return;
            }

            const decoded = jwt(token.token.id_token)
            if (!decoded) {
                next(err, false);
                return;
            }

            const expires = decoded.exp;
            if (!decoded) {
                next(err, false);
                return;
            }

            if (new Date(expires * 1000) < new Date()) {
                next(err, false);
                return;
            }

            const id = token.email;
            if (!id) {
                next(err, false);
                return;
            }

            // Only my email is allowed, fuck every other email
            next(err, id === "michaelrynkiewicz3@gmail.com");
        });
    },

    /**
     * @param {string} code Code that refers to an auth token
     * @param {function(err, ApiKey)} next Function that is called when token has been validated
     */
    getApiKeyByCode: (code, next) => {
        db.getApiKeyByCode(code, (err, apiKey) => {
            if (err) {
                console.log(err);
                next(err, null);
                return;
            }
            next(err, apiKey);
        })
    }
}
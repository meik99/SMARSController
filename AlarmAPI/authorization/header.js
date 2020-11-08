const code = require("./code");

module.exports = {
    authorize: (req, res, next) => {
        if(req.path === '/apps/coffeetogo/api/v1/alarm/health') {
            next();
            return;
        }

        const authorizationHeader = req.header("Authorization")
        if (!authorizationHeader) {
            res
                .status(400)
                .send({
                    code: 400,
                    message: "missing authorization header"
                });
            return;
        }

        const authParts = authorizationHeader.split(" ");
        if (!authParts || authParts.length < 2) {
            res
                .status(400)
                .send({
                    code: 400,
                    message: "invalid authorization header"
                });
            return;
        }

        const authCode = authParts[1];
        code.hasValidToken(authCode, (err, isValid) => {
            if (err) {
                console.log(err);
            }
            if (isValid) {
                next();
            } else {
                code.getApiKeyByCode(authCode, (err, apiKey) => {
                    if (err) {
                        console.log(err);
                    }
                   if (!apiKey) {
                       res
                           .status(401)
                           .send({
                               code: 401,
                               message: "invalid code"
                           });
                       return;
                   }
                   for (let scope of apiKey.scopes) {
                        const scopeParts = scope.split(":");
                        if (scopeParts.length >= 2 && scopeParts[1] !== "") {
                            if(req.method.toLowerCase() === scopeParts[0] && req.path.indexOf(scopeParts[1]) >= 0) {
                                next();
                                return;
                            }
                        }
                   }
                    res
                        .status(401)
                        .send({
                            code: 401,
                            message: `api key is missing scope: ${req.method.toLowerCase()}:<entity>`
                        });
                });
            }
        });
    }
}
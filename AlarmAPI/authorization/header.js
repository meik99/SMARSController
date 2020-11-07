const code = require("./code");

module.exports = {
    authorize: (req, res, next) => {
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
                res
                    .status(401)
                    .send({
                        code: 401,
                        message: "invalid code"
                    });
            }
        });
    }
}
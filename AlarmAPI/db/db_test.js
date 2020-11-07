const assert = require("assert");
const db = require("./db")

describe("Database", () => {
   describe("getTokenForCode", () => {
      it("should connect to the database", (done) => {
          db.getTokenByCode("some-code", (err, token) => {
              assert.strictEqual(err, null);
              // Assert token is null since there cannot be a token with code "some-code"
              assert.strictEqual(token, null);
              done();
          });
      });
    });
});
const code = require("./code");
const assert = require("assert");

describe("code", () => {
   it("should have hasValidToken function", done => {
       code.hasValidToken("some-code", (err, isValid) => {
           assert.strictEqual(err, null);
           assert.strictEqual(isValid, false);
       });
   });
});
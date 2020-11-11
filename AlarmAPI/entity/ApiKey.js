module.exports = class ApiKey {
    constructor(args) {
        if(args) {
            if(args._id) {
                this._id = args._id;
            }
            if (args.scopes) {
                this.scopes = args.scopes;
            } else {
                this.scopes = [];
            }
        }
    }
}
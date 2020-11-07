module.exports = class Alarm {
    constructor(args) {
        if (args) {
            if (args._id) {
                this._id = args._id;
            }
            if (args._rev) {
                this._rev = args._rev;
            }
            if (args.hour) {
                this.hour = args.hour;
            } else {
                this.hour = 0;
            }
            if (args.minute) {
                this.minute = args.minute;
            } else {
                this.minute = 0;
            }
            if (args.days) {
                /*valid values for args.days:
                * []
                * ["mon"]
                * ["mon", "tue", "wed"]
                * ["tue", "fri"]
                * ["mon", "tue", "wed", "thu", "fri", "sat", "sun"]
                * */

                this.days = args.days;
            } else {
                this.days = [];
            }
        }
    }
}
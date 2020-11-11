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
                const hour = parseInt(args.hour);
                this.hour = isNaN(hour) ? -1 : hour;
            } else {
                this.hour = 0;
            }
            if (args.minute) {
                const minute = parseInt(args.minute);
                this.minute = isNaN(minute) ? -1 : minute;
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
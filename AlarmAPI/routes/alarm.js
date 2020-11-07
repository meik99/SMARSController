const express = require('express');
const db = require("../db/db");
const Alarm = require("../entity/Alarm");
const router = express.Router();

/* GET home page. */
router.get('/', function(req, res) {
    db.getAlarms((err, result) => {
        if (err) {
            res
                .status(500)
                .send({
                    code: 500,
                    message: "could not query alarms"
                });
        } else {
            res.send(result);
        }
    });
});

router.post("/", (req, res) => {
    if (!req.body) {
        res
            .status(400)
            .send({
                code: 400,
                message: "missing request body"
            });
        return;
    }

    db.postAlarm(new Alarm(req.body), (err, result) => {
        if (err) {
            res
                .status(500)
                .send({
                    code: 500,
                    message: "could not insert alarm"
                });
        } else {
            res.send(result);
        }
    });
});

router.put("/", (req, res) => {
    if (!req.body) {
        res
            .status(400)
            .send({
                code: 400,
                message: "missing request body"
            });
        return;
    }

    const alarm = new Alarm(req.body);
    if (!alarm._id) {
        res
            .status(400)
            .send({
                code: 400,
                message: "missing id from alarm"
            });
        return;
    }

    db.putAlarm(alarm, (err, result) => {
        if (err) {
            res
                .status(500)
                .send({
                    code: 500,
                    message: "could not update alarm"
                });
        } else {
            res.send(result);
        }
    });
});

router.delete("/:id", (req, res) => {
   if (!req.params || !req.params.id) {
       res
           .status(400)
           .send({
               code: 400,
               message: "query params is missing id"
           });
       return
   }

   db.deleteAlarm(req.params.id, (err, result) => {
        if (err) {
            res
                .status(500)
                .send({
                    code: 500,
                    message: "error deleting alarm"
                });
        } else {
            res.send(result);
        }
   });
});

module.exports = router;

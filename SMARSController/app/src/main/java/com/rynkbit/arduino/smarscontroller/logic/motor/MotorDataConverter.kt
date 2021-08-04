package com.rynkbit.arduino.smarscontroller.logic.motor

import kotlin.math.roundToInt

class MotorDataConverter(private val angle: Int, private val strength: Int) {
    private val MAX_SPEED = 255

    /*
    * Angle - Actuation
    * 90 - L: 1, R: 1
    * 0 / 360 - L: -1, R: 1
    * 270 - L: -1, R: -1
    * 180 - L: 1, R: -1
    * */

    private val leftMotorActuation: Double
        get() {
            return when(angle) {
                in 0..90 -> calculateLinearMotorActuation(angle, k=1, b=-1)
                in 90..180 -> 1.0
                in 180..270 -> calculateLinearMotorActuation(angle, k=-1, b=5)
                in 270..360 -> -1.0
                else -> 0.0
            }
        }

    private val rightMotorActuation: Double
        get() {
            return when(angle) {
                in 0..90 -> 1.0
                in 90..180 -> calculateLinearMotorActuation(angle, k=-1, b=3)
                in 180..270 -> -1.0
                in 270..360 -> calculateLinearMotorActuation(angle, k=1, b=-7)
                else -> 0.0
            }
        }

    private fun calculateLinearMotorActuation(angle: Int, k: Int, b: Int): Double {
        return angle * (k*(1/45.0)) + b
    }

    val leftMotorSpeed: Int
        get() {
            return (strength * (MAX_SPEED / 100.0) * leftMotorActuation).roundToInt()
        }

    val rightMotorSpeed: Int
        get() {
            return (strength * (MAX_SPEED / 100.0) * rightMotorActuation).roundToInt()
        }
}
'use strict';

const TabletClient = require('./tablets/TabletClient');

const client = new TabletClient('http://localhost:8080/');

const yellow = str => `\x1b[0;33m${str}\x1b[0m`;

(async () => {
    console.log(yellow('=== Scenario 1 - non-existing tablet ==='));
    try {
        const telemetry = await client.getTelemetry(0);
        console.log(telemetry);
    } catch (error) {
        console.error(error);
    }

    console.log(yellow('=== Scenario 2 - existing tablet ==='));
    try {
        const telemetry = await client.getTelemetry(1);
        console.log(telemetry);
    } catch (error) {
        console.error(error);
    }

    console.log(yellow('=== Scenario 3 - inserting ==='));
    try {
        const response = await client.setTelemetry(1, [
            {
                battery: 50,
                deviceTime: new Date(Date.now() - 10000),
                current_video: 'Funny dogs'
            },
            {
                battery: 25,
                deviceTime: new Date(),
                current_video: 'Funny cats'
            }
        ]);
        console.dir({ response });
    } catch (error) {
        console.error(error);
    }

    console.log(yellow('=== Scenario 4 - inserting before 10s passed ==='));
    try {
        const response = await client.setTelemetry(1, [
            {
                battery: 10,
                deviceTime: new Date(Date.now() - 10000),
                current_video: 'How to Setup Go server'
            }
        ]);
        console.dir({ response });
    } catch (error) {
        console.error(error);
    }

    console.log(yellow('=== Scenario 5 - reading inserted data ==='));
    try {
        const telemetry = await client.getTelemetry(1);
        console.log(telemetry);
    } catch (error) {
        console.error(error);
    }
})()

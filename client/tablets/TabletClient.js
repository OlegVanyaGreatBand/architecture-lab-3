'use strict';

const HttpJsonClient = require('../common/HttpJsonClient');

module.exports = class TabletClient extends HttpJsonClient {
    constructor(baseUrl) {
        super(baseUrl);
    }

    getTelemetry(id) {
        return this.get(`${this.baseUrl}/${id}`);
    }

    setTelemetry(telemetry) {
        return this.post(this.baseUrl, telemetry);
    }
}

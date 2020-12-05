'use strict';

const HttpJsonClient = require('../common/HttpJsonClient');

module.exports = class TabletClient extends HttpJsonClient {
    constructor(baseUrl) {
        super(baseUrl);
    }

    getTelemetry(id) {
        return this.get(id);
    }

    setTelemetry(tabletId, telemetry) {
        return this.post(this.baseUrl, {
            id: tabletId,
            telemetry
        });
    }
}

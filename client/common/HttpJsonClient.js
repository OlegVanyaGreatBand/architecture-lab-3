'use strict';

const http = require('http');

module.exports = class HttpJsonClient {
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    get(path) {
        return this.#request(this.baseUrl + path, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });
    }

    post(path, data) {
        return this.#request(path, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
        }, JSON.stringify(data));
    }

    async #request(url, options, body) {
        return new Promise((resolve, reject) => {
            const req = http.request(url, options, res => {
                const { statusCode } = res;
                if (statusCode !== 200) {
                    let data = '';
                    res.setEncoding('utf8');
                    res.on('data', chunk => {
                        data += chunk;
                    });
                    res.on('end', () => {
                        try {
                            reject({ code: statusCode, body: JSON.parse(data) });
                        } catch (e) {
                            reject(e);
                        }
                    });

                    return;
                }
                let data = '';
                res.setEncoding('utf8');
                res.on('data', chunk => {
                    data += chunk;
                });
                res.on('end', () => {
                    try {
                        resolve(JSON.parse(data));
                    } catch (e) {
                        reject(e);
                    }
                });
            });

            if (body) {
                req.write(body);
            }

            req.end();
        });
    }
}
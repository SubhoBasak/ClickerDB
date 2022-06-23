const net = require("net");

class StickerClient {
  #socket;
  #result;

  constructor() {
    this.#socket = new net.Socket();
  }

  connect(host = "127.0.0.1", port = 3778) {
    this.#socket.connect(port, host, () =>
      this.#socket.on("data", (data) => (this.#result = data))
    );
  }

  query(q) {
    if (!q.endsWith(";")) q += ";";
    this.#socket.write(q);
  }

  resp() {
    return this.#result.toString();
  }
}

module.exports = StickerClient;

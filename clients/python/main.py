from mimetypes import init
import socket

from traitlets import Bool


class StickerClient:
    def __init__(self) -> None:
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def connect(self, host: str = '127.0.0.1', port: int = 3778) -> None:
        self.socket.connect((host, port))

    def query(self, q: str, buf: int = 1024) -> str:
        if not q.endswith(';'):
            q += ';'
        self.socket.send(q.encode())
        return self.socket.recv(buf).decode()

    # CRUD operations
    def read(self, k: str) -> str:
        pass

    def write(self, k: str, t: str, v) -> bool:
        pass

    def put(self, k: str, v) -> bool:
        pass

    def delete(self, k: str) -> None:
        pass

    def all(self) -> str:
        pass

    def allType(self, t: str) -> str:
        pass

    def clear(self) -> None:
        pass

    # Arithmetic operations
    def add(self, k: str, v: str):
        pass

    def sub(self, k: str, v: str):
        pass

    def mul(self, k: str, v: str):
        pass

    def div(self, k: str, v: str):
        pass

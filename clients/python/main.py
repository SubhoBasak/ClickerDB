import socket


class StickerException(Exception):
    def __init__(self, *args: object) -> None:
        super().__init__(*args)


class StickerClient:
    def __init__(self) -> None:
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def __parseResponse(self, o: str) -> str:
        if o.startswith('E'):
            raise StickerException(o[2:])
        elif o.startswith('O'):
            return o

    def __checkKey(self, k: str) -> bool:
        return False

    def connect(self, host: str = '127.0.0.1', port: int = 3778) -> None:
        self.socket.connect((host, port))

    def query(self, q: str, buf: int = 1024) -> str:
        if not q.endswith(';'):
            q += ';'
        self.socket.send(q.encode())
        return self.__parseResponse()

    # CRUD operations
    def read(self, k: str, buf: int = 1024) -> str:
        self.socket.send(f'R {k};'.encode())
        return self.__parseResponse()

    def write(self, k: str, t: str, v, buf: int = 1024) -> bool:
        self.socket.send(f'W {k} {t} {v};'.encode())
        return self.__parseResponse()

    def put(self, k: str, v, buf: int = 1024) -> bool:
        self.socket.send(f'P {k} {v};'.encode())
        return self.__parseResponse()

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

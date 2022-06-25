import socket


class ClickerClient:
    def __init__(self) -> None:
        '''
            Initialize a new TCP socket to connect with the db server
        '''
        self.__socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def connect(self, host: str = '127.0.0.1', port: int = 3778) -> None:
        '''
            Create a connection with the TCP db server.
        '''
        self.__socket.connect((host, port))

    def close(self):
        self.__socket.close()

    def query(self, q: str, buf: int = 1024) -> str:
        '''
            Run the raw query `q` and return the string with raw
            response data. By default the buffer size is of 1024
            bytes. It can be change using the buf parameter.
        '''
        if not q.endswith(';'):
            # it add a ; at the end of the query if not present
            q += ';'
        self.__socket.send(str(len(q)).rjust(8, '0').encode())
        self.__socket.send(q.encode())
        return self.__socket.recv(buf).decode()

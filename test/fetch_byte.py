import socket
import time

HOST = "127.0.0.1"
PORT = 15001

b = b""


def int_to_255(total) -> list[int]:
    l: list[int] = []
    while total >= 255:
        l.append(total % 255)
        total //= 255
    l.append(total)
    return l


def get_len(content: bytes) -> list[int]:
    l = int_to_255(len(content))
    l.append(255)
    return l


def add_bytes(old: bytes, add: bytes) -> bytes:
    tl = get_len(add)
    return old + bytes(tl) + add


def add_string(old: bytes, content: str) -> bytes:
    t = content.encode()
    return add_bytes(old, t)


b = add_string(b, "fetch_byte")
b = add_string(b, "test.txt")
print(b)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))
s.send(b)
while 1:
    msg = s.recv(1024)
    if msg:
        print(msg)
    time.sleep(1)
s.close()

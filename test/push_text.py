import socket
from util import add_string, add_bytes, int_to_255

HOST = "127.0.0.1"
PORT = 15001

b = b""


b = add_string(b, "text")
b = add_string(b, "分撒发声法是否会撒谎福安市法华寺的的撒的话分撒分撒")
# print(b)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))
s.send(b)
s.close()

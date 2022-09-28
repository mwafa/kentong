# %%
import random
import socket

def getServer(port=39000, password = "SERVER_UDP"):
    client_port = random.randint(30000, 40000)
    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    # Enable broadcasting mode
    client.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
    client.bind(("", client_port))
    client.sendto(password.encode(), ("<broadcast>", port))
    _, addr = client.recvfrom(1024)
    client.close()
    return addr[0]

getServer()

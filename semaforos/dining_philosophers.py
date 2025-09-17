import threading

FILOSOFOS = 5
GARFOS = 5

def filosofo(id, garfo_esquerdo, garfo_direito):
    if id != 0:
        while True:
            print(f"{id} - senta")
            garfo_direito.acquire()
            print(f"{id} - pegou direito")
            garfo_esquerdo.acquire()
            print(f"{id} - pegou esquerdo")
            print(f"{id} - come")
            garfo_direito.release()
            garfo_esquerdo.release()
            print(f"{id} - levanta e pensa")
    else:
        while True:
            print(f"{id} - senta")
            garfo_esquerdo.acquire()
            print(f"{id} - pegou esquerdo")            
            garfo_direito.acquire()
            print(f"{id} - pegou direito")
            print(f"{id} - come")            
            garfo_esquerdo.release()
            garfo_direito.release()
            print(f"{id} - levanta e pensa")

lista_garfos = [threading.Semaphore(1) for i in range(GARFOS)]
for id in range(FILOSOFOS):
    go = threading.Thread(target=filosofo, args=(id, lista_garfos[id], lista_garfos[(id + 1) % GARFOS]))
    go.start()

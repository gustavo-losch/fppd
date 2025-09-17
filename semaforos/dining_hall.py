import threading
import time
import random

eating = 0
readyToLeave = 0

mutex = threading.Semaphore(1)
okToLeave = threading.Semaphore(0)

# Número de estudantes
N = 6

def student(student_id):
    global eating, readyToLeave

    # Simula pegando comida
    print(f"O estudante {student_id} está pegando comida.")
    time.sleep(random.uniform(0.1, 0.5))

    # Entrou no dining hall
    mutex.acquire()
    eating += 1
    print(f"O estudante {student_id} começou a comer. Eating count = {eating}, ReadyToLeave = {readyToLeave}")

    # 2 comendo + 1 esperando (sair)
    if eating == 2 and readyToLeave == 1:
        okToLeave.release()
        readyToLeave -= 1
        print(f"O estudante {student_id} liberou. Eating = {eating}, ReadyToLeave = {readyToLeave}")
    mutex.release()

    # Comendo
    time.sleep(random.uniform(0.5, 1.5))

    # Terminou de comer
    mutex.acquire()
    eating -= 1
    readyToLeave += 1
    print(f"O estudante {student_id} terminou de comer. Eating = {eating}, ReadyToLeave = {readyToLeave}")

    if eating == 1 and readyToLeave == 1:
        # 1 comendo e 1 esperando (esperar)
        mutex.release()
        print(f"O estudante {student_id} está esperando. Eating = {eating}, ReadyToLeave = {readyToLeave}")
        okToLeave.acquire()
    elif eating == 0 and readyToLeave == 2:
        # 2 esperando (sair)
        okToLeave.release()
        readyToLeave -= 2
        print(f"O estudante {student_id} sinaliza: dois estudantes podem sair juntos.")
        mutex.release()
    else:
        # Pronto para sair
        readyToLeave -= 1
        print(f"O estudante {student_id} saiu normalmente.")
        mutex.release()

    print(f"O estudante {student_id} saiu. Eating = {eating}, ReadyToLeave = {readyToLeave}")

def main():
    threads = [threading.Thread(target=student, args=(i,)) for i in range(N)]

    for t in threads:
        t.start()
        time.sleep(random.uniform(0.1, 0.4))  # Atrasar entrada

    for t in threads:
        t.join()

    print("Todos os estudantes sairam.")

if __name__ == "__main__":
    main()

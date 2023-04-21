from random import randint
from asciimatics.screen import Screen
from time import sleep

def demo(screen: Screen):
    screen.print_at('Zen Br', 0, 0)
    while True:
        ev = screen.get_key()
        if ev in (ord('Q'), ord('q')):
            return
        screen.refresh()
Screen.wrapper(demo)
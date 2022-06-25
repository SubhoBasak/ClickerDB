from main import *

c = ClickerClient()
c.connect()

print(c.query("W INT x 123"))
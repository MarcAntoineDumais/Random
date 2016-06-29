import sys
from random import randint

total = 0
for i in range(1, len(sys.argv)):
	raw = sys.argv[i].upper()
	if raw == "CHAR" or raw == "CHARACTER":
		for j in range(6):
			print("best 3 of 4d6: ", end="")
			subtotal = 0
			lowest = 6
			for k in range(4):
				n = randint(1,6)
				print(str(n), end="")
				if k < 3:
					print(", ", end="")
				subtotal = subtotal + n
				if n < lowest:
					lowest = n
			subtotal = subtotal - lowest
			print("  [" + str(subtotal) + "]")
			total = total + subtotal
		break
	numbers = raw.split("D")
	if len(numbers) != 2:
		break
	try:
		amount = int(numbers[0])
		dice = int(numbers[1])
	except ValueError:
		break
	if amount < 1 or dice < 1:
		break
	
	subtotal = 0
	print(str(amount) + "d" + str(dice) + ": ", end="")
	for j in range(amount):
		n = randint(1,dice)
		subtotal = subtotal + n
		print(str(n), end="")
		if j < amount - 1:
			print(", ", end="")
	print("  [" + str(subtotal) + "]")
	total = total + subtotal
	
print("total: " + str(total))
		
# DrunkTeenPattiGame
The festival season is here and you realise it's hard to play all those complicated card games when you're drunk.
You decide to create a simple luck-based game for people to play when they have limited motor and sensory control.

Basic Rules:
- Use a standard deck of cards (no Joker).
- Each player is dealt only three cards.
- 'A' is considered to have a number value of 1.
- 'A' is considered the top card in a face-off. So the order is A > K > Q > J > 10...2

Victory:
- A trail (three cards of the same number) is the highest possible combination.
- The next highest is a sequence (numbers in order, e.g., 4,5,6. A is considered to have a value of 1).
- The next highest is a pair of cards (e.g.: two Kings or two 10s).
- If all else fails, the top card (by number value wins).
- If the top card has the same value, each of the tied players draws a single card from the deck until a winner is found.
- Only the newly drawn cards are compared to decide a tie. The top card wins a tie.
- For now the suit (spades/hearts etc...), does not matter.

Sample :
1. Simulate a game between 4 players.
2. Randomly deal them cards from a deck.
3. Determine the winner.


```shell script
___  ___  __  ___         __        __                __           __         ___  ___  ___          __       ___ ___
 |  |__  /__'  |     \ / /  \ |  | |__)    |    |  | /  ' |__/    |__) \ /     |  |__  |__  |\ |    |__)  /\   |   |  |
 |  |___ .__/  |      |  \__/ \__/ |  \    |___ \__/ \__, |  \    |__)  |      |  |___ |___ | \|    |    /~~\  |   |  |

      ___          __              ___     __  ___       __  ___  ___  __
|\ | |__  |  |    / _'  /\   |\/| |__     /__'  |   /\  |__)  |  |__  |  \
| \| |___ |/\|    \__> /~~\  |  | |___    .__/  |  /~~\ |  \  |  |___ |__/

__________________________________________________________________________

Enter How many Player going to play this Game. : 4

Enter Player 1 Name : Rajan
Rajan (Player ID : G1-P0) Susseccfully Joined
Enter Player 2 Name : Naveen
Naveen (Player ID : G1-P1) Susseccfully Joined
Enter Player 3 Name : Anuj
Anuj (Player ID : G1-P2) Susseccfully Joined
Enter Player 4 Name : Pawan
Pawan (Player ID : G1-P3) Susseccfully Joined
====================================================

Name : Rajan (Player ID : G1-P0)
	.--------.	.--------.	.--------.
	| 2.--.  |	| 5.--.  |	| 7.--.  |
	|  | ♠ | |	|  | ♠ | |	|  | ♦ | |
	|  | ♠ | |	|  | ♠ | |	|  | ♦ | |
	|  '--'2 |	|  '--'5 |	|  '--'7 |
	`--------`	`--------`	`--------`

-------------------------------------------------

Name : Naveen (Player ID : G1-P1)
	.--------.	.--------.	.--------.
	| 9.--.  |	| J.--.  |	| A.--.  |
	|  | ♦ | |	|  | ♣ | |	|  | ♠ | |
	|  | ♦ | |	|  | ♣ | |	|  | ♠ | |
	|  '--'9 |	|  '--'J |	|  '--'A |
	`--------`	`--------`	`--------`

-------------------------------------------------

Name : Anuj (Player ID : G1-P2)
	.--------.	.--------.	.--------.
	| 2.--.  |	| 8.--.  |	| K.--.  |
	|  | ♦ | |	|  | ♦ | |	|  | ♦ | |
	|  | ♦ | |	|  | ♦ | |	|  | ♦ | |
	|  '--'2 |	|  '--'8 |	|  '--'K |
	`--------`	`--------`	`--------`

-------------------------------------------------

Name : Pawan (Player ID : G1-P3) (Winner)
	.--------.	.--------.	.--------.
	| 10.--.  |	| 10.--.  |	| 10.--.  |
	|  | ♥ | |	|  | ♠ | |	|  | ♦ | |
	|  | ♥ | |	|  | ♠ | |	|  | ♦ | |
	|  '--'10 |	|  '--'10 |	|  '--'10 |
	`--------`	`--------`	`--------`

-------------------------------------------------

====================================================
Do you want to continue this this Game, for NO enter 0
```

###Steps to run
1. Clone Repo to your go work Space
`git clone git@github.com:rajankumar549/DrunkTeenPattiGame.git` 

2. Build Binarry 
`go build play.go`

3. Run Binnary
`./play`

4. just run sample binnary in repo

```shell script ./play```


Author : Rajan Kumar 

Email : rajankumar549@gmail.com

Mob : +91-95401-52552
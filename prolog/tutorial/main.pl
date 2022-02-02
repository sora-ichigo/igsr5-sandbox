hi :- write('Hello World!').

grandchild(X, Z) :- child(X, Y), child(Y, Z).

child(nobuyasu, ieyasu).
child(hideyasu, ieyasu).
child(hidetada, ieyasu).
child(tadateru, ieyasu).
child(yoshinao, ieyasu).
child(yorinobu, ieyasu).
child(yorihusa, ieyasu).

child(iemitsu, hidetada).
child(tadanaga, hidetada).
child(masako, hidetada).
child(masayuki, hidetada).
child(mitsusada, yorinobu).

square(X, Y) :- Y is X * X.

fact(0, 1).
fact(N, F) :-
  N > 0,
  N1 is N - 1,
  fact(N1, F1),
  F is N * F1.

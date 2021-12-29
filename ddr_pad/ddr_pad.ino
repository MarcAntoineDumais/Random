#include <Keyboard.h>

// Edit the keys to change which key each pin corresponds to
int keys[4] = {KEY_LEFT_ARROW, KEY_UP_ARROW, KEY_DOWN_ARROW, KEY_RIGHT_ARROW};
int pins[4] = {A0, A1, A2, A3};
bool debug = false;
int lowThreshold[4] = {600, 700, 300, 700};
int highThreshold[4] = {625, 725, 325, 725};

int vals[4];
bool pressed[4];
char buff[50];

void setup() {
  Keyboard.begin();
  if (debug) {
    Serial.begin(9600);
  }
}

void loop() {
  for (int i = 0; i < 4; i++) {
    vals[i] = readKey(pins[i]);
    if (!pressed[i] && vals[i] > highThreshold[i]) {
      pressed[i] = true;
      Keyboard.press(keys[i]);
    }
    if (pressed[i] && vals[i] < lowThreshold[i]) {
      pressed[i] = false;
      Keyboard.release(keys[i]);
    }
  }

  if (debug) {
    printKeys();
  }
  delay(2);
}

int readKey(int pin) {
  analogRead(pin);
  //delay(1);
  return analogRead(pin);
}

void printKeys() {
  for (int i = 0; i < 4; i++) {
    sprintf(buff, "key %d: %4d    ", keys[i], vals[i]);
    Serial.print(buff);
  }
  Serial.println("");
}

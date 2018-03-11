void setup (){
  Serial.begin (9600);

  pinMode (LED_BUILTIN, OUTPUT);
}

void loop (){
  if (Serial.available ()){
    Serial.println (Serial.read ());

    digitalWrite (LED_BUILTIN, HIGH);
    delay (1000);
    digitalWrite (LED_BUILTIN, LOW);
  }
}

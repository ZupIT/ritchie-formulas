#!/bin/sh

#In sh for receive inputs of CLI use: $NAME, $COFFEE_TYPE and $DELIVERY for this exemple

echo "Preparing your coffee $NAME ....."
sleep 1
echo "......"
sleep 1
echo "......"
sleep 1
echo "......"
sleep 1
if [ "$DELIVERY" = true ] ; then
  echo "Your $COFFEE_TYPE coffee is ready, enjoy your trip"
else
  echo "Your $COFFEE_TYPE coffee is ready, have a seat and enjoy your drink"
fi




cd the-final-cl-test/mystery
ID_interview=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $ID_interview 
head interviews/interview-$ID_interview
#MAIN_SUSPECT = "Dartey Henv"
echo $MAIN_SUSPECT
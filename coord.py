# coord.py der nächste Versuch
# Sebastian Kind 23. Oktoboer 2015

from math import *
from tkinter import *

class GUI(object):
    def __init__ (self, ptrack):
        self.track = ptrack

        self.window = Tk()
        self.window.title = ("Coordinates")
        self.window.geometry("400x150")
        self.canvas = Canvas(master=self.window)
        self.canvas.place (x=0, y=0, width=400, height=400)
        self.userInputLat = StringVar()
        self.userInputLong = StringVar()
        self.stringDistanceValue = StringVar ()
        self.stringCoordNumber = StringVar()

        self.labelLat = Label (self.window, text="Latitude").grid(row=0)
        self.labelLong = Label (master=self.window, text="Longitude").grid(row=1, column=0)
        self.entryLat = Entry (self.window, textvariable=self.userInputLat).grid(row=0, column=1)
        self.entryLong = Entry (self.window, textvariable=self.userInputLong).grid(row=1, column=1)

        self.buttonQuit = Button (self.window, text="Quit", command=self.window.quit).grid(row=6, column=3)
        self.buttonNext = Button (self.window, text="next & calculate", command=self.buttonNext).grid(row=2, column=1)
        self.buttonReset = Button (self.window, text="reset track", command=self.buttonReset).grid(row=2, column=0)

        self.labelCoordinates = Label (self.window, text="number of coordinates").grid(row=3, column=0)
        self.labelCoordNumber = Label (self.window, textvariable=self.stringCoordNumber).grid(row=4, column=0)
        self.labelDistance = Label (self.window, text="distance in km").grid(row=3, column=1)
        self.labelDistanceValue = Label (self.window, textvariable=self.stringDistanceValue).grid(row=4, column=1)
        
    
    def buttonNext(self):
        print(self.userInputLat.get(), self.userInputLong.get())
        self.track.appendCoordinates((float)(self.userInputLat.get()), (float)(self.userInputLong.get()))
        self.track.calcTrackDistance()
        self.stringDistanceValue.set((str)(self.track.getTrackDistance()))
        self.stringCoordNumber.set((str)(self.track.getCoordCount()))
        self.userInputLat.set("")
        self.userInputLong.set("")

    def buttonReset(self):
        self.track.setCoordinates([])
        self.track.getCoordCount()
        self.userInputLat.set("")
        self.userInputLong.set("")
        self.stringDistanceValue.set("")
        self.stringCoordNumber.set("")
        

class Track(object):
    def __init__ (self, plistCoordinates):
        self.listCoordinates = plistCoordinates
        self.trackDistance = 0;
        self.coordCount = len(self.listCoordinates)

    def calcDistance (self, alat, along, blat, blong):
        pi = 3.14159265359 #...
        distance = 0
        angle = 0
        deltaLong = 0
        cosDeltaLambda = 0

        a_sin = sin (alat*(pi/180))
        b_sin = sin (blat*(pi/180))
        a_cos = cos (alat*(pi/180))
        b_cos = cos (blat*(pi/180))

        if ((along < 0.0) or (blong < 0.0)):
            if (along < 0.0):
               along *= -1
            if (blong < 0.0):
                blong *= -1
            deltaLong = along + blong

        else:
            deltaLong = along - blong

        cosDeltaLambda = cos (deltaLong*(pi/180))
        angle = acos( a_sin * b_sin + a_cos * b_cos * cosDeltaLambda)

        distance = 2 * pi * 6371 * ((angle * 180/pi)/360)
        return distance
    
    def calcTrackDistance (self):
        self.trackDistance = 0
        i = 0
        while (i+1) < len(self.listCoordinates):
            alat = self.listCoordinates [i] [0]
            along = self.listCoordinates [i] [1]
            blat = self.listCoordinates [i+1] [0]
            blong = self.listCoordinates [i+1] [1]
            i+=1

            self.trackDistance += self.calcDistance(alat, along, blat, blong)




    def getTrackDistance (self):
        return self.trackDistance
    
    def getCoordCount(self):
        self.coordCount = len(self.listCoordinates)
        return self.coordCount

    def getListCoordinates(self):
        return self.listCoordinates

    def insertCoordinate(self, n, alat, along):
        self.listCoordinates = self.listCoordinates[:n] +  [[alat, along]] + self.listCoordinates[n:]
        self.calcTrackDistance()
        self.coordCount+=1

    def appendCoordinates(self, alat, along):
        self.listCoordinates += [[alat, along]]
        self.calcTrackDistance()
        self.coordCount+=1

    def setCoordinates(self, plistCoordinates):
        self.listCoordinates = plistCoordinates
        self.coordCount = len(self.listCoordinates)

    

###### Program ######

myCoords = []
myTrack = Track (myCoords)

myTrack.calcTrackDistance()
print(myTrack.getTrackDistance())

gui = GUI (myTrack)
gui.window.mainloop()

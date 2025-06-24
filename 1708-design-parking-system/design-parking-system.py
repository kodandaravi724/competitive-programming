class ParkingSystem:

    def __init__(self, big: int, medium: int, small: int):
        self.spacesRemaining = [big, medium, small] 

    def addCar(self, carType: int) -> bool:
        self.spacesRemaining[carType-1] -= 1
        return True if self.spacesRemaining[carType-1] >= 0 else False


# Your ParkingSystem object will be instantiated and called as such:
# obj = ParkingSystem(big, medium, small)
# param_1 = obj.addCar(carType)
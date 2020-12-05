slice :: Int -> Int -> [a] -> [a]
slice from to xs = take (to - from + 1) (drop from xs)

maxElement :: [Int] -> Int -> Int
maxElement [] currentMax = currentMax
maxElement (x:xs) currentMax
    | x > currentMax = maxElement xs x
    | otherwise      = maxElement xs currentMax

getPositionID :: String -> Int
getPositionID problemStr =
    (findLine (slice 0 6 problemStr) 0 127) * 8 + (findColumn (slice 7 9 problemStr) 0 7)

findColumn :: String -> Int -> Int -> Int
findColumn [] _ maxValue = maxValue
findColumn (x:xs) minValue maxValue
    | x == 'L' = findColumn xs minValue ((minValue+maxValue) `div` 2)
    | x == 'R' = findColumn xs ((minValue+maxValue) `div` 2) maxValue

findLine :: String -> Int -> Int -> Int
findLine [] _ maxValue = maxValue
findLine (x:xs) minValue maxValue
    | x == 'F' = findLine xs minValue ((minValue+maxValue) `div` 2)
    | x == 'B' = findLine xs ((minValue+maxValue) `div` 2) maxValue

main :: IO()
main = do
    fileContent <- readFile "./input.txt"
    print ("Highest seat ID " ++ show (maxElement (map (\line -> getPositionID line) (lines fileContent)) 0))

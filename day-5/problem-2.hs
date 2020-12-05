slice :: Int -> Int -> [a] -> [a]
slice from to xs = take (to - from + 1) (drop from xs)

getPositionID :: String -> Int
getPositionID problemStr =
    (findLine (slice 0 6 problemStr) 0 127) * 8 + (findColumn (slice 7 9 problemStr) 0 7)

findColumn :: String -> Int -> Int -> Int
findColumn [] _ maxValue = maxValue
findColumn (x : xs) minValue maxValue
    | x == 'L' = findColumn xs minValue ((minValue + maxValue) `div` 2)
    | x == 'R' = findColumn xs ((minValue + maxValue) `div` 2) maxValue

findLine :: String -> Int -> Int -> Int
findLine [] _ maxValue = maxValue
findLine (x : xs) minValue maxValue
    | x == 'F' = findLine xs minValue ((minValue + maxValue) `div` 2)
    | x == 'B' = findLine xs ((minValue + maxValue) `div` 2) maxValue

quicksort1 :: (Ord a) => [a] -> [a]
quicksort1 [] = []
quicksort1 (x : xs) =
    let smallerSorted = quicksort1 [a | a <- xs, a <= x]
        biggerSorted = quicksort1 [a | a <- xs, a > x]
    in smallerSorted ++ [x] ++ biggerSorted

checkArray :: [Int] -> Int -> Int
checkArray [] _ = 9999
checkArray (x:xs) lastNumber
    | x == lastNumber + 1 = checkArray xs x
    | otherwise = x - 1

main :: IO ()
main = do
    fileContent <- readFile "./input.txt"
    let (idsHead:idsTail) = quicksort1 (map (\line -> getPositionID line) (lines fileContent))
    print ("My seat ID is " ++ show (checkArray (idsHead:idsTail) (idsHead-1)))

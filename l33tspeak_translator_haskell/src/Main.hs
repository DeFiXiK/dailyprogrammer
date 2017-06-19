module Main where

import           Data.Char         (isAlpha, toLower, toUpper)
import           Data.Map          (Map)
import qualified Data.Map          as M
import           Data.Maybe        (fromMaybe)
import           Data.String.Utils (replace)

symbolList :: [(Char, String)]
symbolList =
  [ ('A', "4")
  , ('B', "6")
  , ('E', "3")
  , ('I', "1")
  , ('L', "1")
  , ('M', "(V)")
  , ('N', "(\\)")
  , ('O', "0")
  , ('S', "5")
  , ('T', "7")
  , ('V', "\\/")
  , ('W', "`//")
  ]

symbolMap :: Map Char String
symbolMap = M.fromList symbolList

revSymbolMap :: Map String String
revSymbolMap = M.fromList $ stringToCharList symbolList
  where
    stringToCharList :: [(Char, String)] -> [(String, String)]
    stringToCharList  = map stringToChar
    stringToChar :: (Char, String) -> (String, String)
    stringToChar (one, two) = (two, [toLower one])

inputList :: [String]
inputList = [
  "I am elite.", "Da pain!", "Eye need help!",
  "3Y3 (\\)33d j00 t0 g37 d4 d0c70r.", "1 n33d m4 p1llz!"
  ]

engToLeet :: String -> String
engToLeet = concatMap conv
  where
    conv :: Char -> String
    conv c = fromMaybe [c] (M.lookup (toUpper c) symbolMap)

data Lang = Eng | Leet

wtf :: String -> Lang
wtf str =
  let result = isAlpha (head str)
    in
    if result then Eng else Leet

leetToEng :: String -> String
leetToEng input = foldl (\input (from, to) -> replace from to input) input (M.toList revSymbolMap)

translate :: String -> String
translate input =
  case lang of
    Eng  -> engToLeet input
    Leet -> leetToEng input
  where
    lang = wtf input

main :: IO ()
main = mapM_ (putStrLn . translate) inputList

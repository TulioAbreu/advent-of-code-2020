using System;
using System.Collections.Generic;

struct Bag
{
    public string color;
    public LinkedList<Constraint> constraints;
}

struct Constraint
{
    public int quantity;
    public string color;
}

class Program
{
    const string INPUT_FILEPATH = "../../../input.txt";

    static Constraint parseConstraintStr(string constraintStr)
    {
        string[] constraintChunks = constraintStr.Trim().Split(" ");
        return new Constraint
        {
            quantity = int.Parse(constraintChunks[0].Trim()),
            color = constraintChunks[1] + " " + constraintChunks[2]
        };
    }

    static LinkedList<Constraint> parseConstraintsStr(string constraintsStr)
    {
        if (constraintsStr.Trim() == "no other bags.")
        {
            return new LinkedList<Constraint>();
        }

        LinkedList<Constraint> constraints = new LinkedList<Constraint>();
        foreach (string constraintStr in constraintsStr.Split(","))
        {
            constraints.AddLast(parseConstraintStr(constraintStr.Trim()));
        }
        return constraints;
    }

    static Bag parseBagStr(string bagStr)
    {
        string[] bagColorAndConstraints = bagStr.Split("bags contain");
        Bag bag = new Bag {
            color = bagColorAndConstraints[0].Trim(),
            constraints = parseConstraintsStr(bagColorAndConstraints[1].Trim())
        };
        return bag;
    }

    static LinkedList<Bag> readInputFile(string filepath)
    {
        string fileContent = System.IO.File.ReadAllText(filepath);
        string[] fileLines = fileContent.Split("\n");
        LinkedList<Bag> bags = new LinkedList<Bag>();
        foreach (string line in fileLines)
        {
            Bag bag = parseBagStr(line);
            bags.AddLast(bag);
        }
        return bags;
    }

    static Bag? getBagWithColor(LinkedList<Bag> bags, string targetColor)
    {
        foreach (Bag bag in bags)
        {
            if (bag.color == targetColor)
            {
                return bag;
            }
        }
        return null;
    }

    static bool canContainBag(LinkedList<Bag> bags, Bag currentBag, string targetBag)
    {
        foreach (Constraint constraintBag in currentBag.constraints)
        {
            if (constraintBag.color == targetBag)
            {
                return true;
            }

            Bag? nextCurrentBag = getBagWithColor(bags, constraintBag.color);
            if (!nextCurrentBag.HasValue)
            {
                continue;
            }
            
            if (canContainBag(bags, nextCurrentBag.Value, targetBag))
            {
                return true;
            }
        }
        return false;
    }

    static int getBagCountInside(LinkedList<Bag> bags, Bag currentBag)
    {
        if (currentBag.constraints.Count == 0)
        {
            return 0;
        }

        int sum = 0;
        foreach (Constraint constraintBag in currentBag.constraints)
        {
            Bag? bag = getBagWithColor(bags, constraintBag.color);
            if (!bag.HasValue)
            {
                continue;
            }
            sum += constraintBag.quantity + constraintBag.quantity * getBagCountInside(bags, bag.Value);
        }
        return sum;
    }

    static void Main(string[] args)
    {
        LinkedList<Bag> bags = readInputFile(INPUT_FILEPATH);

        // Part 1
        int counter = 0;
        foreach (Bag bag in bags)
        {
            if (canContainBag(bags, bag, "shiny gold")) counter++;
        }
        Console.WriteLine("Bags that can contain Shiny Gold Bag = " + counter);

        // Part 2
        Bag? shinyGoldBag = getBagWithColor(bags, "shiny gold");
        if (!shinyGoldBag.HasValue)
        {
            Console.Write("Could not find Shiny Gold Bag inside input.");
            return;
        }

        Console.WriteLine("Bags inside Shiny Gold Bag = " + getBagCountInside(bags, shinyGoldBag.Value));
    }
}

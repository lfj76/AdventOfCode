public static class Day1
{
    public static void Solve()
    {
        using StreamReader sr = new("../../inputs/day1.txt");

        List<int> left = new();
        List<int> right = new();

        string? line;
        while ((line = sr.ReadLine()) != null)
        {
            string[] numbers = line.Split(' ', 2, options: StringSplitOptions.TrimEntries);
            left.Add(Int32.Parse(numbers[0]));
            right.Add(Int32.Parse(numbers[1]));
        }

        left.Sort();
        right.Sort();

        List<int> distances = new();

        for (int i = 0; i < left.Count; i++)
        {
            distances.Add(Math.Abs(left[i] - right[i]));
        }

        //Console.WriteLine(distances.Sum());
        int ll = 0;
        int ul = 0;
        int lr = 0;
        int ur = 0;
        int similarity = 0;
        int sim;
        while (ll < left.Count && lr < right.Count)
        {
            while (ul < left.Count && left[ll] == left[ul])
                ul++;
            while (lr + 1 < right.Count && right[lr] < left[ll])
                lr++;
            if (ur < lr)
                ur = lr;
            while (ur < right.Count && right[lr] == right[ur])
                ur++;
            if (left[ll] == right[lr])
            {
                sim = left[ll] * (ul - ll) * (ur - lr);
                //Console.WriteLine($"ll, ul, lr, ur: {ll}, {ul}, {lr}, {ur} => sim: {sim}");
                similarity += sim;
            }
            ll = ul;
            //Console.WriteLine($"ll: {ll}, lr: {lr}");
            //Console.ReadLine();

        }

        Console.WriteLine(similarity);
    }
}

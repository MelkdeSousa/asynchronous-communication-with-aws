namespace Common;

public static class DotEnv
{
    /// <summary>
    /// Loads environment variables from a file and sets them in the current environment.
    /// </summary>
    /// <param name="filePath">The path to the file containing the environment variables.</param>
    public static void Load(string filePath)
    {
        if (!File.Exists(filePath))
            return;

        foreach (var line in File.ReadAllLines(filePath))
        {
            var parts = line.Split(
                '=',
                StringSplitOptions.RemoveEmptyEntries);

            if (parts.Length != 2)
                continue;

            Environment.SetEnvironmentVariable(parts[0], parts[1]);
        }
    }

}
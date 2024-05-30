namespace Database;

public class User
{
    public Guid UserId { get; set; }

    public List<Point> Points { get; set; }

    public User()
    {
        Points = new List<Point>();
    }
}
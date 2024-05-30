namespace Database;

public class Point
{
    public Guid PointID { get; set; }
    public DateTime DateTime { get; set; }

    public Guid UserID { get; set; }
    public virtual User? User { get; set; }

    public string? Type { get; set; }
}

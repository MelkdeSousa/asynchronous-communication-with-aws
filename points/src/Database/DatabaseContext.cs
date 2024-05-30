using Microsoft.EntityFrameworkCore;

namespace Database;

public class DatabaseContext : DbContext
{
    public DbSet<Point> Points { get; set; }
    public DbSet<User> Users { get; set; }

    public DatabaseContext() { }

    public DatabaseContext(DbContextOptions<DatabaseContext> options)
                : base(options) { }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
           => optionsBuilder.UseNpgsql(Environment.GetEnvironmentVariable("ASPNETCORE_DATABASE_URL")!);
}

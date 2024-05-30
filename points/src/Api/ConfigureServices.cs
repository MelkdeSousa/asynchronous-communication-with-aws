using Database;
using Microsoft.EntityFrameworkCore;

namespace Microsoft.Extensions.DependencyInjection;

public static class ConfigureServices
{
    public static IServiceCollection AddDatabase(this IServiceCollection services)
    {
        return services.AddDbContext<DatabaseContext>(options => options.UseNpgsql(Environment.GetEnvironmentVariable("ASPNETCORE_DATABASE_URL")!));
    }
}

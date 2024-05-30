using Database;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace Api.Controllers;

[Route("api/points")]
[ApiController]
public class PointController : ControllerBase
{
        private readonly DatabaseContext _context;


    public PointController(DatabaseContext context)
    {

        _context = context;
    }

    [HttpGet(Name = "GetPoints")]
    public async Task<IEnumerable<Point>> Get()
    {
        var points = await _context.Points.ToListAsync();

        return points;
    }
}


using Api.HttpDTOs;
using Database;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace Api.Controllers;

[Route("api/points", Name = "Points")]
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

    [HttpPost("{userId}", Name = "CreatePoint")]
    public async Task<ActionResult<Point>> Post(
        [FromBody] CreatePointBodyRequest input,
        [FromRoute] Guid userId
    )
    {
        var user = await _context.Users.FindAsync(userId);
        if (user == null)
        {
            return NotFound("User not found");
        }

        var point = await _context.Points.AddAsync(
            new Point
            {
                DateTime = input.DateTime,
                Type = input.Type,
                UserID = userId,
                PointID = Guid.NewGuid(),
                User = user
            }
        );
        await _context.SaveChangesAsync();

        return Created("/api/points/" + point.Entity.PointID, point.Entity);
    }
}

# Hotel Manager

*[Read this in Russian / Читать на русском →](./README_ru.md)*

---

1. [Description](#description)
2. [Use Cases](#use-cases)
3. [Development Iterations](#development-iterations)

---

## Description

A hotel management and control service.

Written in Golang using the Telegram API. PostgreSQL is used as the database.

---

## Use Cases

The service has a role-based model. The following actors are involved:

1. **Owner**. The owner is the main administrator of the system and has full access to managing all aspects of the hotel’s operations.

   - Receives reporting information about the hotel.
   - Manages staff, which includes:
     adding/removing employees (managers, receptionists, housekeepers),
     assigning managers responsible for different aspects of hotel operations.
   - Can manage pricing policies for rooms and other hotel services.

2. **Manager-Administrator**. The manager-administrator handles the day-to-day management and coordination of the hotel’s operations.

   - Controls room reservations and monitors timely check-ins and check-outs.
   - Maintains contact with the owner and provides necessary reporting information.
   - Processes customer complaints and requests, resolving issues as they arise.
   - Monitors the condition of hotel premises and initiates repairs when necessary.
   - Can add and modify information about room and service availability in the system.

3. **Receptionist**. The receptionist interacts directly with hotel guests and serves as their primary point of contact.

   - Handles guest registration during check-in and check-out.
   - Keeps records of bookings and checks the availability of free rooms.
   - Processes guest requests and wishes during their stay (additional services, information, etc.).
   - Reports any issues or violations related to guests to the manager-administrator.

4. **Housekeeper**. The housekeeper is responsible for cleanliness and order in guest rooms and common areas.

   - Performs daily room cleaning according to the schedule.
   - Informs the receptionist or manager-administrator about any malfunctions or damage in rooms.
   - Maintains cleanliness in the hotel’s public areas (lobbies, corridors, etc.).
   - Reports completion of room cleaning through the system.

---

## Development Iterations

### Iteration 1. Basic functionality and stability.

In the first iteration, the project foundation was laid and basic functionality was added in the form of 12 commands covering rooms, clients, employees, and room occupancy.

### Iteration 2. Functionality expansion.

In the second iteration, a role-based model was implemented using middleware. On the front-end side (Telegram), more convenient buttons were added for interacting with the system. Package restructuring and code refactoring were also performed.

### Iteration 3. Full feature completion and public release.

All features were finalized and the system’s usability (UI/UX) was improved. System behavior was enhanced, texts were corrected, and additional refactoring and code improvements were carried out.


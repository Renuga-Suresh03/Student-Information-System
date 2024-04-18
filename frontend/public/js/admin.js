/*// JavaScript for admin pages

// Function to handle admin login form submission
function handleAdminLogin() {
    var adminId = document.getElementById('adminId').value;
    var password = document.getElementById('password').value;

    // Perform client-side validation if needed

    // Send login request to backend API
    fetch('/api/admin/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            admin_id: adminId,
            password: password
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to admin home page or display success message
        console.log('Admin login successful:', data);
        window.location.href = '/admin/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Admin login failed:', error);
        alert('Admin login failed. Please try again.');
    });
}

// Function to handle navigation to admin home page
function navigateToAdminHome() {
    window.location.href = '/admin/home.html';
}

// Function to handle navigation to admin profile page
function navigateToAdminProfile() {
    window.location.href = '/admin/profile.html';
}

// Function to handle navigation to admin marks page
function navigateToAdminMarks() {
    window.location.href = '/admin/marks.html';
}

// Function to handle navigation to admin attendance page
function navigateToAdminAttendance() {
    window.location.href = '/admin/attendance.html';
}

// Function to handle adding marks
function addMarks() {
    // Implement functionality to add marks
    console.log('Functionality to add marks');
}

// Function to handle adding attendance
function addAttendance() {
    // Implement functionality to add attendance
    console.log('Functionality to add attendance');
}*/

// admin.js login

// Function to handle admin login form submission
function handleAdminLogin() {
    var adminId = document.getElementById('adminId').value;
    var password = document.getElementById('password').value;

    // Perform client-side validation if needed

    // Send login request to backend API
    fetch('/api/admin/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            admin_id: adminId,
            password: password
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to admin home page or display success message
        console.log('Admin login successful:', data);
        window.location.href = '/admin/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Admin login failed:', error);
        alert('Admin login failed. Please try again.');
    });
}

//js for attendance

function addAttendance() {
    var attendanceData = [];

    // Loop through each row in the table
    var tableRows = document.querySelectorAll('#attendanceTable tbody tr');
    tableRows.forEach(function(row) {
        var regNo = row.cells[0].innerText;
        var name = row.cells[1].innerText;
        var dates = [];
        var status = [];
        var studentId = row.dataset.studentId;
        var presentCount = 0; // Track the number of present days

        // Loop through each date column to get status
        for (var i = 2; i <= 8; i++) { // Loop through 7 date columns
            var selectElement = row.cells[i].querySelector('select');
            var selectedValue = selectElement.value;
            status.push(selectedValue);
            if (selectedValue === 'P') {
                presentCount++; // Increment present count if status is 'P'
            }
            dates.push(selectElement.selectedIndex); // Save selected index of the dropdown
        }

        // Initial percentage calculation
        var totalCount = status.length;
        var initialPercentage = totalCount > 0 ? Math.round((presentCount / totalCount) * 100) : '-';

        attendanceData.push({
            regNo: regNo,
            name: name,
            dates: dates,
            status: status,
            studentId: studentId,
            percentage: initialPercentage
        });
    });

    // Function to update the percentages in the table
    function updatePercentages() {
        tableRows.forEach(function(row) {
            var status = [];
            var presentCount = 0;
            // Loop through each date column to get status
            for (var i = 2; i <= 8; i++) { // Loop through 7 date columns
                var selectElement = row.cells[i].querySelector('select');
                var selectedValue = selectElement.value;
                status.push(selectedValue);
                if (selectedValue === 'P') {
                    presentCount++;
                }
            }
            var totalCount = status.length;
            var percentage = totalCount > 0 ? Math.round((presentCount / totalCount) * 100) : 0; // Calculate the percentage without appending '%'
            row.cells[9].innerText = percentage + '%'; // Update the percentage column

            // Highlight the percentage cell if it goes below 75%
            if (percentage < 75) {
                row.cells[9].classList.add('low-percentage');
            } else {
                row.cells[9].classList.remove('low-percentage');
            }
        });
    }

    // Call the function to update percentages
    updatePercentages();

    // Send attendance data to backend or perform further processing
    console.log('Attendance data:', attendanceData);
}

/*
//js for marks

// Function to calculate and update total, percentage, and rank
function updateCalculations() {
    var rows = document.querySelectorAll("#marksTable tbody tr");

    // Update total marks and calculate percentage for each row
    rows.forEach(function(row, index) {
        var subjects = row.querySelectorAll('.subjectInput');
        var total = 0;

        // Calculate total marks
        subjects.forEach(function(subject) {
            total += parseInt(subject.value) || 0;
        });

        var totalCell = row.cells[5];
        var percentageCell = row.cells[6];
        var rankCell = row.cells[7];

        totalCell.textContent = total;
        percentageCell.textContent = ((total / (subjects.length * 100)) * 100).toFixed(2) + "%";
    });

    // Update rank based on percentage
    var sortedRows = Array.from(rows).sort((a, b) => {
        var percentageA = parseFloat(a.cells[6].textContent);
        var percentageB = parseFloat(b.cells[6].textContent);
        return percentageB - percentageA;
    });

    var previousRank = 0;
    var previousPercentage = 0;
    sortedRows.forEach(function(sortedRow, sortedIndex) {
        var currentPercentage = parseFloat(sortedRow.cells[6].textContent);
        var currentRank = sortedIndex + 1;

        // Update rank only if the percentage is different from the previous row
        if (currentPercentage !== previousPercentage) {
            rankCell = sortedRow.cells[7];
            rankCell.textContent = currentRank;
            previousRank = currentRank;
        } else {
            // If the percentage is the same as the previous row, use the previous rank
            rankCell = sortedRow.cells[7];
            rankCell.textContent = previousRank;
        }
        previousPercentage = currentPercentage;
    });
}

// Event listener for save button
document.getElementById("saveBtn").addEventListener('click', function() {
    // Call updateCalculations to update total, percentage, and rank
    updateCalculations();
});



// Function to update result based on percentage
// Function to update result based on percentage
function updateCalculations() {
    var rows = document.querySelectorAll("#marksTable tbody tr");

    // Update total marks and calculate percentage for each row
    rows.forEach(function(row, index) {
        var subjects = row.querySelectorAll('.subjectInput');
        var total = 0;
        var isAbsent = false;

        // Calculate total marks and check for absence or marks less than 50
        subjects.forEach(function(subject) {
            var mark = parseInt(subject.value) || 0;
            if (subject.value === '-' || mark < 50) {
                isAbsent = true;
            }
            total += mark;
        });

        var totalCell = row.cells[5];
        var percentageCell = row.cells[6];
        var rankCell = row.cells[7];

        totalCell.textContent = isAbsent ? '-' : total;
        percentageCell.textContent = isAbsent ? '-' : ((total / (subjects.length * 100)) * 100).toFixed(2) + "%";
        rankCell.textContent = isAbsent ? '-' : ''; // Clear rank for absent students
    });

    // Update rank based on percentage for non-absent students
    var sortedRows = Array.from(rows).filter(row => row.cells[7].textContent !== '-').sort((a, b) => {
        var percentageA = parseFloat(a.cells[6].textContent);
        var percentageB = parseFloat(b.cells[6].textContent);
        return percentageB - percentageA;
    });

    var previousRank = 0;
    var previousPercentage = 0;
    sortedRows.forEach(function(sortedRow, sortedIndex) {
        var currentPercentage = parseFloat(sortedRow.cells[6].textContent);
        var currentRank = sortedIndex + 1;

        // Update rank only if the percentage is different from the previous row
        if (currentPercentage !== previousPercentage) {
            var rankCell = sortedRow.cells[7];
            rankCell.textContent = currentRank;
            previousRank = currentRank;
        } else {
            // If the percentage is the same as the previous row, use the previous rank
            var rankCell = sortedRow.cells[7];
            rankCell.textContent = previousRank;
        }
        previousPercentage = currentPercentage;
    });
}
*/

//js for marks
// Function to calculate and update total, percentage, and rank
function updateCalculations() {
    var rows = document.querySelectorAll("#marksTable tbody tr");
    var totalAll = 0;

    // Update total marks and calculate percentage for each row
    rows.forEach(function(row, index) {
        var subjects = row.querySelectorAll('.subjectInput');
        var total = 0;
        var isAbsent = false;

        // Calculate total marks
        subjects.forEach(function(subject, subjectIndex) {
            var mark = parseInt(subject.value) || 0;
            if (subject.value === '-' || mark < 50) {
                isAbsent = true;
                if (mark < 50) {
                    subject.style.backgroundColor = 'lightcoral'; // Highlight marks less than 50 in light red color
                }
                return; // Skip further calculations if mark is absent or less than 50
            } else {
                subject.style.backgroundColor = ''; // Remove any previous background color
            }
            total += mark;
        });

        var totalCell = row.cells[5];
        var percentageCell = row.cells[6];
        var rankCell = row.cells[7];

        // Update total cell
        totalCell.textContent = isAbsent ? '-' : total;

        // Calculate percentage only if not absent and total is above 50
        if (!isAbsent && total >= 50) {
            var percentage = ((total / (subjects.length * 100)) * 100).toFixed(2) + "%";
            percentageCell.textContent = percentage;
        } else {
            percentageCell.textContent = '-';
            rankCell.textContent = ''; // Clear rank if absent or marks less than 50
            return; // Skip generating rank for this row
        }

        // Generate rank for non-absent students
        rankCell.textContent = generateRank(rows, index, total);

        // Add total to overall total
        totalAll += total;
    });

    // Update total for all rows
    var totalAllCell = document.getElementById("totalAll");
    totalAllCell.textContent = totalAll;
}

// Function to generate rank for non-absent students
function generateRank(rows, currentIndex, currentTotal) {
    var sortedRows = Array.from(rows).filter((row, index) => {
        if (index === currentIndex) return false; // Exclude current row
        var subjects = row.querySelectorAll('.subjectInput');
        for (var subject of subjects) {
            if (subject.value === '-' || parseInt(subject.value) < 50) {
                return false; // Exclude rows with '-' or marks less than 50
            }
        }
        return true; // Include rows with valid marks
    }).sort((a, b) => {
        var totalA = parseInt(a.cells[5].textContent) || 0;
        var totalB = parseInt(b.cells[5].textContent) || 0;
        return totalB - totalA;
    });

    var rank = 1;
    sortedRows.forEach(function(sortedRow, i) {
        var currentRowTotal = parseInt(sortedRow.cells[5].textContent) || 0;
        if (currentRowTotal > currentTotal) {
            rank = i + 2; // Increment rank for each row with higher total marks
        }
    });
    return rank;
}

// Event listener for save button
document.getElementById("saveBtn").addEventListener('click', function() {
    // Call updateCalculations to update total, percentage, and rank
    updateCalculations();
});

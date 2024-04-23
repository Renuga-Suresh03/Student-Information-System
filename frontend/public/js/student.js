// JavaScript for student pages

// Function to handle student login form submission
function handleStudentLogin() {
    var regNo = document.getElementById('regNo').value;
    var dob = document.getElementById('dob').value;

    // Send login request to backend API
    fetch('/api/student/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            reg_no: regNo,
            dob: dob
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to student home page or display success message
        console.log('Student login successful:', data);
        window.location.href = '/student/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Student login failed:', error);
        alert('Student login failed. Please try again.');
    });
}

// Function to handle navigation to student home page
function navigateToStudentHome() {
    window.location.href = '/student/home.html';
}

// Function to handle navigation to student profile page
function navigateToStudentProfile() {
    window.location.href = '/student/profile.html';
}

// Function to handle navigation to student marks page for each exam
function navigateToStudentMarks(examType) {
    window.location.href = `/student/assessment${examType}.html`;
}

// Function to handle navigation to student attendance page
function navigateToStudentAttendance() {
    window.location.href = '/student/attendance.html';
}

// Function to fetch and display assessment marks for a student
function displayAssessmentMarks(examType) {
    // Fetch assessment marks for the specified exam type from backend API
    fetch(`/api/student/assessment-marks?examType=${examType}`)
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to fetch assessment marks');
        }
        return response.json();
    })
    .then(data => {
        // Process and display assessment marks
        console.log(`Assessment marks for ${examType}:`, data);
        // Implement logic to display assessment marks on the page
    })
    .catch(error => {
        console.error('Failed to fetch assessment marks:', error);
        // Display error message to user
    });
}

// Function to fetch and display attendance records for a student
function displayAttendanceRecords() {
    // Fetch attendance records for the student from backend API
    fetch('/api/student/attendance')
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to fetch attendance records');
        }
        return response.json();
    })
    .then(data => {
        // Process and display attendance records
        console.log('Attendance records:', data);
        // Implement logic to display attendance records on the page
    })
    .catch(error => {
        console.error('Failed to fetch attendance records:', error);
        // Display error message to user
    });
}

// Function to fetch and display student marks data dynamically for each exam
function fetchAndDisplayMarksData(examType) {
    // Fetch marks data for the student for the specified exam from backend API
    fetch(`/api/student/marks?examType=${examType}`)
    .then(response => {
        if (!response.ok) {
            throw new Error(`Failed to fetch marks data for ${examType}`);
        }
        return response.json();
    })
    .then(data => {
        // Process and display marks data
        console.log(`Marks data for ${examType}:`, data);
        // Implement logic to display marks data on the page
    })
    .catch(error => {
        console.error(`Failed to fetch marks data for ${examType}:`, error);
        // Display error message to user
    });
}

// Event listener for loading marks data when the page is loaded
document.addEventListener("DOMContentLoaded", function() {
    // Fetch and display marks data for each exam when the page is loaded
    fetchAndDisplayMarksData('assessment1');
    fetchAndDisplayMarksData('assessment2');
    fetchAndDisplayMarksData('model');
});

// Event listener for loading attendance data when the page is loaded
document.addEventListener("DOMContentLoaded", function() {
    displayAttendanceRecords();
});


// Sample code to dynamically generate table rows for attendance
// Sample code to dynamically generate table rows for attendance and display attendance percentage with a progress bar
document.addEventListener("DOMContentLoaded", function() {
    var attendanceData = [
        { date: "2024-04-01", status: "Present" },
        { date: "2024-04-02", status: "Absent" },
        { date: "2024-04-03", status: "Present" },
        { date: "2024-04-04", status: "Present" },
        { date: "2024-04-05", status: "Absent" },
        { date: "2024-04-06", status: "Present" },
        { date: "2024-04-07", status: "Present" }
    ];

    var tableBody = document.querySelector("#attendanceTable tbody");
    attendanceData.forEach(function(item) {
        var row = document.createElement("tr");
        var dateCell = document.createElement("td");
        var statusCell = document.createElement("td");
        dateCell.textContent = item.date;
        statusCell.textContent = item.status;
        row.appendChild(dateCell);
        row.appendChild(statusCell);
        if (item.status === "Present") {
            statusCell.classList.add("present");
        } else {
            statusCell.classList.add("absent");
        }
        tableBody.appendChild(row);
    });

    // Calculate and display attendance percentage with progress bar
    var totalDays = attendanceData.length;
    var presentDays = attendanceData.filter(item => item.status === "Present").length;
    var attendancePercentage = (presentDays / totalDays * 100).toFixed(2);
    document.getElementById("attendancePercentage").textContent = "Attendance Percentage: " + attendancePercentage + "%";
    var graphBar = document.createElement("div");
    graphBar.classList.add("graph-bar");
    graphBar.style.width = attendancePercentage + "%";
    document.querySelector(".graph").appendChild(graphBar);
});


// Sample marks data for three students and three subjects
// Sample marks data for three subjects
const marksData = [
    { subjectCode: "101", subjectName: "Mathematics", marks: 60 },
    { subjectCode: "102", subjectName: "Science", marks: 30 },
    { subjectCode: "103", subjectName: "English", marks: 60 }
];

// Function to populate the marks table
function populateMarksTable() {
    const tableBody = document.querySelector("#marksTable tbody");
    let totalMarks = 0;

    marksData.forEach((subject, index) => {
        const row = tableBody.insertRow();
        const slNoCell = row.insertCell(0);
        const subjectCodeCell = row.insertCell(1);
        const subjectNameCell = row.insertCell(2);
        const marksCell = row.insertCell(3);
        const percentageCell = row.insertCell(4);
        const statusCell = row.insertCell(5);

        slNoCell.textContent = index + 1;
        subjectCodeCell.textContent = subject.subjectCode;
        subjectNameCell.textContent = subject.subjectName;
        marksCell.textContent = subject.marks;
        percentageCell.textContent = calculatePercentage(subject.marks) + "%";
        totalMarks += subject.marks;
        statusCell.textContent = subject.marks > 50 ? "Pass" : "Fail";
        if (subject.marks <= 50) {
            statusCell.classList.add("fail");
        }
    });

    const totalElement = document.getElementById("total");
    totalElement.textContent = totalMarks;

    const finalPercentageBox = document.getElementById("finalPercentageBox");
    finalPercentageBox.textContent = "Final Percentage: " + (totalMarks / (marksData.length * 100)) * 100 + "%";
}

// Function to calculate percentage
function calculatePercentage(marks) {
    return (marks / 100) * 100;
}

// Call the function to populate marks table
populateMarksTable();
